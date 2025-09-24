// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetJobConfigResponseBody
	GetRequestId() *string
}

type DeleteDatasetJobConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DA869D1B-035A-43B2-ACC1-C56681BD9FAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetJobConfigResponseBody) SetRequestId(v string) *DeleteDatasetJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
