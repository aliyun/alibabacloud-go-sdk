// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMmsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *StopMmsJobResponseBody
	GetData() *int64
	SetRequestId(v string) *StopMmsJobResponseBody
	GetRequestId() *string
}

type StopMmsJobResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 8023D058-62B7-5C49-8EB6-AD9BA7942BC5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopMmsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopMmsJobResponseBody) GetData() *int64 {
	return s.Data
}

func (s *StopMmsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMmsJobResponseBody) SetData(v int64) *StopMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *StopMmsJobResponseBody) SetRequestId(v string) *StopMmsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMmsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
