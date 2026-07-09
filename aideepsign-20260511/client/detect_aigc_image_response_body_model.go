// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectAigcImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*DetectAigcImageResponseBodyBody) *DetectAigcImageResponseBody
	GetBody() []*DetectAigcImageResponseBodyBody
	SetCode(v string) *DetectAigcImageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DetectAigcImageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DetectAigcImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetectAigcImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetectAigcImageResponseBody
	GetSuccess() *bool
}

type DetectAigcImageResponseBody struct {
	// The list of AIGC detection result labels.
	Body []*DetectAigcImageResponseBodyBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Repeated"`
	// The business error code. The value `OK` is returned if the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value `200` is returned if the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. The value `success` is returned if the request was successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectAigcImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectAigcImageResponseBody) GoString() string {
	return s.String()
}

func (s *DetectAigcImageResponseBody) GetBody() []*DetectAigcImageResponseBodyBody {
	return s.Body
}

func (s *DetectAigcImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetectAigcImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DetectAigcImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetectAigcImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectAigcImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetectAigcImageResponseBody) SetBody(v []*DetectAigcImageResponseBodyBody) *DetectAigcImageResponseBody {
	s.Body = v
	return s
}

func (s *DetectAigcImageResponseBody) SetCode(v string) *DetectAigcImageResponseBody {
	s.Code = &v
	return s
}

func (s *DetectAigcImageResponseBody) SetHttpStatusCode(v int32) *DetectAigcImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetectAigcImageResponseBody) SetMessage(v string) *DetectAigcImageResponseBody {
	s.Message = &v
	return s
}

func (s *DetectAigcImageResponseBody) SetRequestId(v string) *DetectAigcImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectAigcImageResponseBody) SetSuccess(v bool) *DetectAigcImageResponseBody {
	s.Success = &v
	return s
}

func (s *DetectAigcImageResponseBody) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetectAigcImageResponseBodyBody struct {
	// The confidence level. Value range: 0 to 1. A higher value indicates a higher probability.
	//
	// example:
	//
	// 0.51
	Confidence *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The detection label. Valid values:
	//
	// - `ai_generated`: AI-generated.
	//
	// - `non_ai_generated`: Not AI-generated.
	//
	// example:
	//
	// ai_generated
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DetectAigcImageResponseBodyBody) String() string {
	return dara.Prettify(s)
}

func (s DetectAigcImageResponseBodyBody) GoString() string {
	return s.String()
}

func (s *DetectAigcImageResponseBodyBody) GetConfidence() *string {
	return s.Confidence
}

func (s *DetectAigcImageResponseBodyBody) GetLabel() *string {
	return s.Label
}

func (s *DetectAigcImageResponseBodyBody) SetConfidence(v string) *DetectAigcImageResponseBodyBody {
	s.Confidence = &v
	return s
}

func (s *DetectAigcImageResponseBodyBody) SetLabel(v string) *DetectAigcImageResponseBodyBody {
	s.Label = &v
	return s
}

func (s *DetectAigcImageResponseBodyBody) Validate() error {
	return dara.Validate(s)
}
