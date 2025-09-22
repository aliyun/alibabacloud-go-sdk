// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckNFCBindPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *CheckNFCBindPopRequest
	GetActivityId() *int64
	SetNfcId(v string) *CheckNFCBindPopRequest
	GetNfcId() *string
}

type CheckNFCBindPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	NfcId *string `json:"NfcId,omitempty" xml:"NfcId,omitempty"`
}

func (s CheckNFCBindPopRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckNFCBindPopRequest) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *CheckNFCBindPopRequest) GetNfcId() *string {
	return s.NfcId
}

func (s *CheckNFCBindPopRequest) SetActivityId(v int64) *CheckNFCBindPopRequest {
	s.ActivityId = &v
	return s
}

func (s *CheckNFCBindPopRequest) SetNfcId(v string) *CheckNFCBindPopRequest {
	s.NfcId = &v
	return s
}

func (s *CheckNFCBindPopRequest) Validate() error {
	return dara.Validate(s)
}
