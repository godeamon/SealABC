/*
 * Copyright 2020 The SealABC Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package uidData

import (
	"github.com/SealSC/SealABC/dataStructure/enum"
	"github.com/SealSC/SealABC/metadata/seal"
)

var UIDKeyTypes struct {
	SelfSignature   enum.Element
	OracleSignature enum.Element
}

type UIDKey struct {
	KeyType  int
	KeyData  []byte
	KeyProof []byte
}

type UniversalIdentificationData struct {
	Identification string
	Keys           []UIDKey
}

type UniversalIdentification struct {
	UniversalIdentificationData
	Seal seal.Entity
}
