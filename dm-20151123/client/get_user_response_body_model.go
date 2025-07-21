// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetUserResponseBodyData) *GetUserResponseBody
	GetData() *GetUserResponseBodyData
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
}

type GetUserResponseBody struct {
	// Returned Content
	Data *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetData() *GetUserResponseBodyData {
	return s.Data
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyData struct {
	// Whether EventBridge is enabled
	//
	// example:
	//
	// true
	EnableEventbridge *bool `json:"EnableEventbridge,omitempty" xml:"EnableEventbridge,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) GetEnableEventbridge() *bool {
	return s.EnableEventbridge
}

func (s *GetUserResponseBodyData) SetEnableEventbridge(v bool) *GetUserResponseBodyData {
	s.EnableEventbridge = &v
	return s
}

func (s *GetUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
