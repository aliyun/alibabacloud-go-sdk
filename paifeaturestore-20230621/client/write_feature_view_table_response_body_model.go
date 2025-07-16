// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWriteFeatureViewTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WriteFeatureViewTableResponseBody
	GetRequestId() *string
	SetTaskId(v string) *WriteFeatureViewTableResponseBody
	GetTaskId() *string
}

type WriteFeatureViewTableResponseBody struct {
	// example:
	//
	// 0C89F5E1-7F24-5EEC-9F05-508A39278CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s WriteFeatureViewTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WriteFeatureViewTableResponseBody) GoString() string {
	return s.String()
}

func (s *WriteFeatureViewTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WriteFeatureViewTableResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *WriteFeatureViewTableResponseBody) SetRequestId(v string) *WriteFeatureViewTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *WriteFeatureViewTableResponseBody) SetTaskId(v string) *WriteFeatureViewTableResponseBody {
	s.TaskId = &v
	return s
}

func (s *WriteFeatureViewTableResponseBody) Validate() error {
	return dara.Validate(s)
}
