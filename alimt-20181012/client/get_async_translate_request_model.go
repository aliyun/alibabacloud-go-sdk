// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAsyncTranslateRequest
	GetJobId() *string
}

type GetAsyncTranslateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 77056ab7-7be1-4c2a-91a1-f20f63894048
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAsyncTranslateRequest) SetJobId(v string) *GetAsyncTranslateRequest {
	s.JobId = &v
	return s
}

func (s *GetAsyncTranslateRequest) Validate() error {
	return dara.Validate(s)
}
