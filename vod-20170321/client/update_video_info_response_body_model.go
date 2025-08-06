// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVideoInfoResponseBody
	GetRequestId() *string
}

type UpdateVideoInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVideoInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoInfoResponseBody) SetRequestId(v string) *UpdateVideoInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
