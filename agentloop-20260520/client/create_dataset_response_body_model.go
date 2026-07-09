// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
}

type CreateDatasetResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EB27D183-8F6C-5C5A-A6A3-E0508AF54F78
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
