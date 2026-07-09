// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDatasetResponseBody
	GetRequestId() *string
}

type UpdateDatasetResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BC5B6F62-6FA2-57FC-8285-99753BD34B6D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
