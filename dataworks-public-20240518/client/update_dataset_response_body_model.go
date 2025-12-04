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
	SetSuccess(v bool) *UpdateDatasetResponseBody
	GetSuccess() *bool
}

type UpdateDatasetResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *UpdateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetSuccess(v bool) *UpdateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
