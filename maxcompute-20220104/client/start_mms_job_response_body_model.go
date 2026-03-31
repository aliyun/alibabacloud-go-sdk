// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMmsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *StartMmsJobResponseBody
	GetData() *int64
	SetRequestId(v string) *StartMmsJobResponseBody
	GetRequestId() *string
}

type StartMmsJobResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 5CA6292A-E301-5CD8-B4E2-AF060F99147B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartMmsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartMmsJobResponseBody) GetData() *int64 {
	return s.Data
}

func (s *StartMmsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMmsJobResponseBody) SetData(v int64) *StartMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *StartMmsJobResponseBody) SetRequestId(v string) *StartMmsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMmsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
