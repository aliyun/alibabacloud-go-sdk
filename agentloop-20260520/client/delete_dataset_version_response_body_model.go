// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetVersionResponseBody
	GetRequestId() *string
}

type DeleteDatasetVersionResponseBody struct {
	// The unique identifier of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D17DE39E-6C62-50E3-9EB7-FDE41BB0D43D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetVersionResponseBody) SetRequestId(v string) *DeleteDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
