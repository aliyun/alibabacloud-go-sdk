// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomImageNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomImageNameResponseBody
	GetRequestId() *string
}

type UpdateCustomImageNameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 552B7EED-D434-511F-B838-29EA4E906034
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomImageNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomImageNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomImageNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomImageNameResponseBody) SetRequestId(v string) *UpdateCustomImageNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomImageNameResponseBody) Validate() error {
	return dara.Validate(s)
}
