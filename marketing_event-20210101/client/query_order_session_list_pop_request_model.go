// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderSessionListPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *QueryOrderSessionListPopRequest
	GetActivityId() *int64
	SetNfcId(v string) *QueryOrderSessionListPopRequest
	GetNfcId() *string
}

type QueryOrderSessionListPopRequest struct {
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

func (s QueryOrderSessionListPopRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderSessionListPopRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderSessionListPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *QueryOrderSessionListPopRequest) GetNfcId() *string {
	return s.NfcId
}

func (s *QueryOrderSessionListPopRequest) SetActivityId(v int64) *QueryOrderSessionListPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QueryOrderSessionListPopRequest) SetNfcId(v string) *QueryOrderSessionListPopRequest {
	s.NfcId = &v
	return s
}

func (s *QueryOrderSessionListPopRequest) Validate() error {
	return dara.Validate(s)
}
