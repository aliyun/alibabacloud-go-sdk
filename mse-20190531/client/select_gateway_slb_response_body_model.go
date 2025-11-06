// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectGatewaySlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SelectGatewaySlbResponseBody
	GetCode() *int32
	SetData(v []*SelectGatewaySlbResponseBodyData) *SelectGatewaySlbResponseBody
	GetData() []*SelectGatewaySlbResponseBodyData
	SetHttpStatusCode(v int32) *SelectGatewaySlbResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SelectGatewaySlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *SelectGatewaySlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SelectGatewaySlbResponseBody
	GetSuccess() *bool
}

type SelectGatewaySlbResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	Data []*SelectGatewaySlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E00C6D90-A28A-5813-8981-0459AA436F46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SelectGatewaySlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectGatewaySlbResponseBody) GoString() string {
	return s.String()
}

func (s *SelectGatewaySlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SelectGatewaySlbResponseBody) GetData() []*SelectGatewaySlbResponseBodyData {
	return s.Data
}

func (s *SelectGatewaySlbResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SelectGatewaySlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SelectGatewaySlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectGatewaySlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SelectGatewaySlbResponseBody) SetCode(v int32) *SelectGatewaySlbResponseBody {
	s.Code = &v
	return s
}

func (s *SelectGatewaySlbResponseBody) SetData(v []*SelectGatewaySlbResponseBodyData) *SelectGatewaySlbResponseBody {
	s.Data = v
	return s
}

func (s *SelectGatewaySlbResponseBody) SetHttpStatusCode(v int32) *SelectGatewaySlbResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SelectGatewaySlbResponseBody) SetMessage(v string) *SelectGatewaySlbResponseBody {
	s.Message = &v
	return s
}

func (s *SelectGatewaySlbResponseBody) SetRequestId(v string) *SelectGatewaySlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectGatewaySlbResponseBody) SetSuccess(v bool) *SelectGatewaySlbResponseBody {
	s.Success = &v
	return s
}

func (s *SelectGatewaySlbResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SelectGatewaySlbResponseBodyData struct {
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp14lqiw5n96hq2****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The name of the SLB instance.
	//
	// example:
	//
	// MseGatewaySlb-gw
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
}

func (s SelectGatewaySlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SelectGatewaySlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *SelectGatewaySlbResponseBodyData) GetSlbId() *string {
	return s.SlbId
}

func (s *SelectGatewaySlbResponseBodyData) GetSlbName() *string {
	return s.SlbName
}

func (s *SelectGatewaySlbResponseBodyData) SetSlbId(v string) *SelectGatewaySlbResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *SelectGatewaySlbResponseBodyData) SetSlbName(v string) *SelectGatewaySlbResponseBodyData {
	s.SlbName = &v
	return s
}

func (s *SelectGatewaySlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
