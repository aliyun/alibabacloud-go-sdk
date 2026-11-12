// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatasetVersionResponseBody
	GetRequestId() *string
}

type CreateDatasetVersionResponseBody struct {
	// The unique identifier of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D17DE39E-6C62-50E3-9EB7-FDE41BB0D43D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetVersionResponseBody) SetRequestId(v string) *CreateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
