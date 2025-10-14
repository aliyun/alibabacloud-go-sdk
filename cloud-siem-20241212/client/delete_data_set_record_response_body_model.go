// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataSetRecordResponseBody
	GetRequestId() *string
}

type DeleteDataSetRecordResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSetRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSetRecordResponseBody) SetRequestId(v string) *DeleteDataSetRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSetRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
