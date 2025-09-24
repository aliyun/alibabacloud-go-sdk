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
}

type UpdateDatasetVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *UpdateDatasetVersionResponseBody) SetRequestId(v string) *UpdateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
