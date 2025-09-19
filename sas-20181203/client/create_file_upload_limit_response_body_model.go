// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUploadLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFileUploadLimitResponseBody
	GetRequestId() *string
}

type CreateFileUploadLimitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileUploadLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUploadLimitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileUploadLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileUploadLimitResponseBody) SetRequestId(v string) *CreateFileUploadLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileUploadLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
