// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedMonitorStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAdvancedMonitorStateResponseBody
	GetCode() *string
	SetData(v []*GetAdvancedMonitorStateResponseBodyData) *GetAdvancedMonitorStateResponseBody
	GetData() []*GetAdvancedMonitorStateResponseBodyData
	SetMessage(v string) *GetAdvancedMonitorStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAdvancedMonitorStateResponseBody
	GetRequestId() *string
}

type GetAdvancedMonitorStateResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetAdvancedMonitorStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAdvancedMonitorStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedMonitorStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdvancedMonitorStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAdvancedMonitorStateResponseBody) GetData() []*GetAdvancedMonitorStateResponseBodyData {
	return s.Data
}

func (s *GetAdvancedMonitorStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAdvancedMonitorStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdvancedMonitorStateResponseBody) SetCode(v string) *GetAdvancedMonitorStateResponseBody {
	s.Code = &v
	return s
}

func (s *GetAdvancedMonitorStateResponseBody) SetData(v []*GetAdvancedMonitorStateResponseBodyData) *GetAdvancedMonitorStateResponseBody {
	s.Data = v
	return s
}

func (s *GetAdvancedMonitorStateResponseBody) SetMessage(v string) *GetAdvancedMonitorStateResponseBody {
	s.Message = &v
	return s
}

func (s *GetAdvancedMonitorStateResponseBody) SetRequestId(v string) *GetAdvancedMonitorStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdvancedMonitorStateResponseBody) Validate() error {
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

type GetAdvancedMonitorStateResponseBodyData struct {
	// Indicates whether the DPI feature is enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s GetAdvancedMonitorStateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedMonitorStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAdvancedMonitorStateResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetAdvancedMonitorStateResponseBodyData) SetEnable(v bool) *GetAdvancedMonitorStateResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetAdvancedMonitorStateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
