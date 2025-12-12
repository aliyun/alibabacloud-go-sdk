// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityCheckSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualityCheckSchemeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteQualityCheckSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteQualityCheckSchemeResponseBody
	GetMessage() *string
	SetMessages(v *DeleteQualityCheckSchemeResponseBodyMessages) *DeleteQualityCheckSchemeResponseBody
	GetMessages() *DeleteQualityCheckSchemeResponseBodyMessages
	SetRequestId(v string) *DeleteQualityCheckSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityCheckSchemeResponseBody
	GetSuccess() *bool
}

type DeleteQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteQualityCheckSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityCheckSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualityCheckSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityCheckSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualityCheckSchemeResponseBody) GetMessages() *DeleteQualityCheckSchemeResponseBodyMessages {
	return s.Messages
}

func (s *DeleteQualityCheckSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityCheckSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityCheckSchemeResponseBody) SetCode(v string) *DeleteQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *DeleteQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetMessage(v string) *DeleteQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetMessages(v *DeleteQualityCheckSchemeResponseBodyMessages) *DeleteQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetRequestId(v string) *DeleteQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) SetSuccess(v bool) *DeleteQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteQualityCheckSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteQualityCheckSchemeResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityCheckSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteQualityCheckSchemeResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *DeleteQualityCheckSchemeResponseBodyMessages) SetMessage(v []*string) *DeleteQualityCheckSchemeResponseBodyMessages {
	s.Message = v
	return s
}

func (s *DeleteQualityCheckSchemeResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
