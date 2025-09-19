// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileUploadLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *UpdateFileUploadLimitRequest
	GetLimit() *int64
}

type UpdateFileUploadLimitRequest struct {
	// The QPS limit on the files uploaded from the client. Valid values: 100 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s UpdateFileUploadLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileUploadLimitRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileUploadLimitRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *UpdateFileUploadLimitRequest) SetLimit(v int64) *UpdateFileUploadLimitRequest {
	s.Limit = &v
	return s
}

func (s *UpdateFileUploadLimitRequest) Validate() error {
	return dara.Validate(s)
}
