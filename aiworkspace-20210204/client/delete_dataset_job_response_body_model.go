// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetJobResponseBody
	GetRequestId() *string
}

type DeleteDatasetJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetJobResponseBody) SetRequestId(v string) *DeleteDatasetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetJobResponseBody) Validate() error {
	return dara.Validate(s)
}
