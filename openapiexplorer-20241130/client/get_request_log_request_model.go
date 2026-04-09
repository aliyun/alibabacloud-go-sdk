// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogRequestId(v string) *GetRequestLogRequest
	GetLogRequestId() *string
}

type GetRequestLogRequest struct {
	// The request ID returned by the API for which you want to query the log. The value is the universally unique identifiers (UUID) of the API request and must be uppercase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123E4567-E89B-12D3-A456-426614174000
	LogRequestId *string `json:"logRequestId,omitempty" xml:"logRequestId,omitempty"`
}

func (s GetRequestLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogRequest) GoString() string {
	return s.String()
}

func (s *GetRequestLogRequest) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *GetRequestLogRequest) SetLogRequestId(v string) *GetRequestLogRequest {
	s.LogRequestId = &v
	return s
}

func (s *GetRequestLogRequest) Validate() error {
	return dara.Validate(s)
}
