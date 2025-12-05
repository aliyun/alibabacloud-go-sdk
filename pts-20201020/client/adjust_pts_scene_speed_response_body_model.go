// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustPtsSceneSpeedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AdjustPtsSceneSpeedResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AdjustPtsSceneSpeedResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AdjustPtsSceneSpeedResponseBody
	GetMessage() *string
	SetRequestId(v string) *AdjustPtsSceneSpeedResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AdjustPtsSceneSpeedResponseBody
	GetSuccess() *bool
}

type AdjustPtsSceneSpeedResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, no data is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F7D2CE0-XXXX-4143-955A-8E4595AF86A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AdjustPtsSceneSpeedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AdjustPtsSceneSpeedResponseBody) GoString() string {
	return s.String()
}

func (s *AdjustPtsSceneSpeedResponseBody) GetCode() *string {
	return s.Code
}

func (s *AdjustPtsSceneSpeedResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AdjustPtsSceneSpeedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AdjustPtsSceneSpeedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AdjustPtsSceneSpeedResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AdjustPtsSceneSpeedResponseBody) SetCode(v string) *AdjustPtsSceneSpeedResponseBody {
	s.Code = &v
	return s
}

func (s *AdjustPtsSceneSpeedResponseBody) SetHttpStatusCode(v int32) *AdjustPtsSceneSpeedResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AdjustPtsSceneSpeedResponseBody) SetMessage(v string) *AdjustPtsSceneSpeedResponseBody {
	s.Message = &v
	return s
}

func (s *AdjustPtsSceneSpeedResponseBody) SetRequestId(v string) *AdjustPtsSceneSpeedResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdjustPtsSceneSpeedResponseBody) SetSuccess(v bool) *AdjustPtsSceneSpeedResponseBody {
	s.Success = &v
	return s
}

func (s *AdjustPtsSceneSpeedResponseBody) Validate() error {
	return dara.Validate(s)
}
