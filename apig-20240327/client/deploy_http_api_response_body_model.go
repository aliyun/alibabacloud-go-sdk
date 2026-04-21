// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeployHttpApiResponseBody
	GetCode() *string
	SetData(v *DeployHttpApiResponseBodyData) *DeployHttpApiResponseBody
	GetData() *DeployHttpApiResponseBodyData
	SetMessage(v string) *DeployHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeployHttpApiResponseBody
	GetRequestId() *string
}

type DeployHttpApiResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeployHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0C2D1C68-0D93-5561-8EE6-FDB7BF067A30
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeployHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeployHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeployHttpApiResponseBody) GetData() *DeployHttpApiResponseBodyData {
	return s.Data
}

func (s *DeployHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeployHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployHttpApiResponseBody) SetCode(v string) *DeployHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *DeployHttpApiResponseBody) SetData(v *DeployHttpApiResponseBodyData) *DeployHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *DeployHttpApiResponseBody) SetMessage(v string) *DeployHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *DeployHttpApiResponseBody) SetRequestId(v string) *DeployHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployHttpApiResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployHttpApiResponseBodyData struct {
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
}

func (s DeployHttpApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployHttpApiResponseBodyData) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *DeployHttpApiResponseBodyData) SetHttpApiId(v string) *DeployHttpApiResponseBodyData {
	s.HttpApiId = &v
	return s
}

func (s *DeployHttpApiResponseBodyData) Validate() error {
	return dara.Validate(s)
}
