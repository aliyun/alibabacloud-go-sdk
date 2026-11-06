// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataIngestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataIngestionResponseBody
	GetRequestId() *string
}

type DeleteDataIngestionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataIngestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataIngestionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataIngestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataIngestionResponseBody) SetRequestId(v string) *DeleteDataIngestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataIngestionResponseBody) Validate() error {
	return dara.Validate(s)
}
