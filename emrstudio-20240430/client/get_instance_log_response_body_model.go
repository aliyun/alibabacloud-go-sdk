// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetInstanceLogResponseBodyData) *GetInstanceLogResponseBody
	GetData() *GetInstanceLogResponseBodyData
	SetRequestId(v string) *GetInstanceLogResponseBody
	GetRequestId() *string
}

type GetInstanceLogResponseBody struct {
	// example:
	//
	// p-123****
	Data *GetInstanceLogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInstanceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceLogResponseBody) GetData() *GetInstanceLogResponseBodyData {
	return s.Data
}

func (s *GetInstanceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLogResponseBody) SetData(v *GetInstanceLogResponseBodyData) *GetInstanceLogResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceLogResponseBody) SetRequestId(v string) *GetInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceLogResponseBodyData struct {
	// example:
	//
	// 1
	LineNum *int64 `json:"lineNum,omitempty" xml:"lineNum,omitempty"`
	// example:
	//
	// "logs"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetInstanceLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceLogResponseBodyData) GetLineNum() *int64 {
	return s.LineNum
}

func (s *GetInstanceLogResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceLogResponseBodyData) SetLineNum(v int64) *GetInstanceLogResponseBodyData {
	s.LineNum = &v
	return s
}

func (s *GetInstanceLogResponseBodyData) SetMessage(v string) *GetInstanceLogResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetInstanceLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
