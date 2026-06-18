// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLEvaluateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSQLEvaluateTaskResponseBodyData) *CreateSQLEvaluateTaskResponseBody
	GetData() *CreateSQLEvaluateTaskResponseBodyData
	SetMessage(v string) *CreateSQLEvaluateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSQLEvaluateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSQLEvaluateTaskResponseBody
	GetSuccess() *bool
}

type CreateSQLEvaluateTaskResponseBody struct {
	// The returned result.
	Data *CreateSQLEvaluateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. > This parameter is empty when the request succeeds. If the request fails, an error message is returned, such as an error code.
	//
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6352AC16-76BF-5135-B1EA-ED49293526E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSQLEvaluateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLEvaluateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSQLEvaluateTaskResponseBody) GetData() *CreateSQLEvaluateTaskResponseBodyData {
	return s.Data
}

func (s *CreateSQLEvaluateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSQLEvaluateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSQLEvaluateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSQLEvaluateTaskResponseBody) SetData(v *CreateSQLEvaluateTaskResponseBodyData) *CreateSQLEvaluateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateSQLEvaluateTaskResponseBody) SetMessage(v string) *CreateSQLEvaluateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSQLEvaluateTaskResponseBody) SetRequestId(v string) *CreateSQLEvaluateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSQLEvaluateTaskResponseBody) SetSuccess(v bool) *CreateSQLEvaluateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSQLEvaluateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSQLEvaluateTaskResponseBodyData struct {
	// The Slink task ID.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s CreateSQLEvaluateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLEvaluateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSQLEvaluateTaskResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateSQLEvaluateTaskResponseBodyData) SetSlinkTaskId(v string) *CreateSQLEvaluateTaskResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateSQLEvaluateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
