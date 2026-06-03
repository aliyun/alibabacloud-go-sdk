// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigVerifySceneAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfigVerifySceneAppInfoResponseBody
	GetCode() *string
	SetData(v *ConfigVerifySceneAppInfoResponseBodyData) *ConfigVerifySceneAppInfoResponseBody
	GetData() *ConfigVerifySceneAppInfoResponseBodyData
	SetMessage(v string) *ConfigVerifySceneAppInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigVerifySceneAppInfoResponseBody
	GetRequestId() *string
}

type ConfigVerifySceneAppInfoResponseBody struct {
	// example:
	//
	// OK
	Code    *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *ConfigVerifySceneAppInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigVerifySceneAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigVerifySceneAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigVerifySceneAppInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfigVerifySceneAppInfoResponseBody) GetData() *ConfigVerifySceneAppInfoResponseBodyData {
	return s.Data
}

func (s *ConfigVerifySceneAppInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigVerifySceneAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigVerifySceneAppInfoResponseBody) SetCode(v string) *ConfigVerifySceneAppInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBody) SetData(v *ConfigVerifySceneAppInfoResponseBodyData) *ConfigVerifySceneAppInfoResponseBody {
	s.Data = v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBody) SetMessage(v string) *ConfigVerifySceneAppInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBody) SetRequestId(v string) *ConfigVerifySceneAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigVerifySceneAppInfoResponseBodyData struct {
	// example:
	//
	// OK
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigVerifySceneAppInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigVerifySceneAppInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) SetCode(v string) *ConfigVerifySceneAppInfoResponseBodyData {
	s.Code = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) SetMessage(v string) *ConfigVerifySceneAppInfoResponseBodyData {
	s.Message = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) SetRequestId(v string) *ConfigVerifySceneAppInfoResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ConfigVerifySceneAppInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
