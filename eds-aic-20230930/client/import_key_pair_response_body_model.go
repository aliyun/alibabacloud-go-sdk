// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ImportKeyPairResponseBodyData) *ImportKeyPairResponseBody
	GetData() *ImportKeyPairResponseBodyData
	SetRequestId(v string) *ImportKeyPairResponseBody
	GetRequestId() *string
}

type ImportKeyPairResponseBody struct {
	// The object that is returned.
	Data *ImportKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) GetData() *ImportKeyPairResponseBodyData {
	return s.Data
}

func (s *ImportKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportKeyPairResponseBody) SetData(v *ImportKeyPairResponseBodyData) *ImportKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKeyPairResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportKeyPairResponseBodyData struct {
	// The time when the ADB key pair was created.
	//
	// example:
	//
	// 2023-03-05T10:29:22Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3*****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the ADB key pair.
	//
	// example:
	//
	// TestKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s ImportKeyPairResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ImportKeyPairResponseBodyData) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *ImportKeyPairResponseBodyData) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ImportKeyPairResponseBodyData) SetGmtCreated(v string) *ImportKeyPairResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *ImportKeyPairResponseBodyData) SetKeyPairId(v string) *ImportKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *ImportKeyPairResponseBodyData) SetKeyPairName(v string) *ImportKeyPairResponseBodyData {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponseBodyData) Validate() error {
	return dara.Validate(s)
}
