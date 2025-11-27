// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteJobRecordsResponseBody
	GetRequestId() *string
}

type DeleteJobRecordsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobRecordsResponseBody) SetRequestId(v string) *DeleteJobRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
