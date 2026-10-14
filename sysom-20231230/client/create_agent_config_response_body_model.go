// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAgentConfigResponseBody
	GetCode() *string
	SetData(v *CreateAgentConfigResponseBodyData) *CreateAgentConfigResponseBody
	GetData() *CreateAgentConfigResponseBodyData
	SetMessage(v string) *CreateAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAgentConfigResponseBody
	GetRequestId() *string
}

type CreateAgentConfigResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization fails, check the `message` field for the detailed fault message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreateAgentConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error message of the request.
	//
	// example:
	//
	// Requests for llm service failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAgentConfigResponseBody) GetData() *CreateAgentConfigResponseBodyData {
	return s.Data
}

func (s *CreateAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentConfigResponseBody) SetCode(v string) *CreateAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentConfigResponseBody) SetData(v *CreateAgentConfigResponseBodyData) *CreateAgentConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentConfigResponseBody) SetMessage(v string) *CreateAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentConfigResponseBody) SetRequestId(v string) *CreateAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentConfigResponseBodyData struct {
	// The ID of the created component configuration.
	//
	// example:
	//
	// 4d5ecf02-3d01-11f0-8c90-00163e1029af
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
}

func (s CreateAgentConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentConfigResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *CreateAgentConfigResponseBodyData) SetConfigId(v string) *CreateAgentConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *CreateAgentConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
