// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUploadLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *CreateFileUploadLimitRequest
	GetLimit() *int64
}

type CreateFileUploadLimitRequest struct {
	// The QPS limit on the files uploaded from the client. Valid values: 100 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s CreateFileUploadLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUploadLimitRequest) GoString() string {
	return s.String()
}

func (s *CreateFileUploadLimitRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CreateFileUploadLimitRequest) SetLimit(v int64) *CreateFileUploadLimitRequest {
	s.Limit = &v
	return s
}

func (s *CreateFileUploadLimitRequest) Validate() error {
	return dara.Validate(s)
}
