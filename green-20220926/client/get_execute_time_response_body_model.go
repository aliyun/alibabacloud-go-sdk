// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetExecuteTimeResponseBody
	GetData() *string
	SetRequestId(v string) *GetExecuteTimeResponseBody
	GetRequestId() *string
}

type GetExecuteTimeResponseBody struct {
	// example:
	//
	// 02:24:30
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExecuteTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteTimeResponseBody) GetData() *string {
	return s.Data
}

func (s *GetExecuteTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecuteTimeResponseBody) SetData(v string) *GetExecuteTimeResponseBody {
	s.Data = &v
	return s
}

func (s *GetExecuteTimeResponseBody) SetRequestId(v string) *GetExecuteTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecuteTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
