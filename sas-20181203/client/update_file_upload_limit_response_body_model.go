// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileUploadLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileUploadLimitResponseBody
	GetRequestId() *string
}

type UpdateFileUploadLimitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8BF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileUploadLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileUploadLimitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileUploadLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileUploadLimitResponseBody) SetRequestId(v string) *UpdateFileUploadLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileUploadLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
