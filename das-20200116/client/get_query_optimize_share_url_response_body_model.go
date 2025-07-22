// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeShareUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeShareUrlResponseBody
	GetCode() *string
	SetData(v string) *GetQueryOptimizeShareUrlResponseBody
	GetData() *string
	SetMessage(v string) *GetQueryOptimizeShareUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeShareUrlResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeShareUrlResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeShareUrlResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The share URL.
	//
	// example:
	//
	// https://hdm.console.aliyun.com/#/queryOptimize?Keywords=&OnlyOptimizedSql=true&Time=1684771200000&Engine=MySQL&InstanceIds=&Rules=&PageNo=1&PageSize=10&OrderBy=count&Asc=false&SqlIds=&dbNames=&region=cn-china&user=
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeShareUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeShareUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeShareUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeShareUrlResponseBody) GetData() *string {
	return s.Data
}

func (s *GetQueryOptimizeShareUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeShareUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeShareUrlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeShareUrlResponseBody) SetCode(v string) *GetQueryOptimizeShareUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeShareUrlResponseBody) SetData(v string) *GetQueryOptimizeShareUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetQueryOptimizeShareUrlResponseBody) SetMessage(v string) *GetQueryOptimizeShareUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeShareUrlResponseBody) SetRequestId(v string) *GetQueryOptimizeShareUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeShareUrlResponseBody) SetSuccess(v string) *GetQueryOptimizeShareUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeShareUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
