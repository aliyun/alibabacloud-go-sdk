// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionListPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *QuerySessionListPopRequest
	GetActivityId() *int64
	SetNfcId(v string) *QuerySessionListPopRequest
	GetNfcId() *string
}

type QuerySessionListPopRequest struct {
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

func (s QuerySessionListPopRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionListPopRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionListPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *QuerySessionListPopRequest) GetNfcId() *string {
	return s.NfcId
}

func (s *QuerySessionListPopRequest) SetActivityId(v int64) *QuerySessionListPopRequest {
	s.ActivityId = &v
	return s
}

func (s *QuerySessionListPopRequest) SetNfcId(v string) *QuerySessionListPopRequest {
	s.NfcId = &v
	return s
}

func (s *QuerySessionListPopRequest) Validate() error {
	return dara.Validate(s)
}
