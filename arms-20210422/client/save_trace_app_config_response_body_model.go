// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTraceAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SaveTraceAppConfigResponseBody
	GetData() *string
	SetRequestId(v string) *SaveTraceAppConfigResponseBody
	GetRequestId() *string
}

type SaveTraceAppConfigResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveTraceAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTraceAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *SaveTraceAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTraceAppConfigResponseBody) SetData(v string) *SaveTraceAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetRequestId(v string) *SaveTraceAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
