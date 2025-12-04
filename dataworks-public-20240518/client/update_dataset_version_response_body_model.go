// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDatasetVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDatasetVersionResponseBody
	GetSuccess() *bool
}

type UpdateDatasetVersionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A090666C-74FB-5629-ABFC-2FE99DD55XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDatasetVersionResponseBody) SetRequestId(v string) *UpdateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetVersionResponseBody) SetSuccess(v bool) *UpdateDatasetVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
