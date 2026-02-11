// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTraceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteTraceAppResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteTraceAppResponseBody
	GetRequestId() *string
}

type DeleteTraceAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTraceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteTraceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTraceAppResponseBody) SetData(v string) *DeleteTraceAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetRequestId(v string) *DeleteTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTraceAppResponseBody) Validate() error {
	return dara.Validate(s)
}
