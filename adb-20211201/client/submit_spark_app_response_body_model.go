// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSparkAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitSparkAppResponseBodyData) *SubmitSparkAppResponseBody
	GetData() *SubmitSparkAppResponseBodyData
	SetRequestId(v string) *SubmitSparkAppResponseBody
	GetRequestId() *string
}

type SubmitSparkAppResponseBody struct {
	// The returned data.
	Data *SubmitSparkAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSparkAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkAppResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppResponseBody) GetData() *SubmitSparkAppResponseBodyData {
	return s.Data
}

func (s *SubmitSparkAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSparkAppResponseBody) SetData(v *SubmitSparkAppResponseBodyData) *SubmitSparkAppResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSparkAppResponseBody) SetRequestId(v string) *SubmitSparkAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSparkAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSparkAppResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// s202204132018hzprec1ac61a000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// TestApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The alert message returned for the operation, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	//
	// example:
	//
	// Insufficient resources.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution state of the application. Valid values:
	//
	// 	- **SUBMITTED**
	//
	// 	- **STARTING**
	//
	// 	- **RUNNING**
	//
	// 	- **FAILING**
	//
	// 	- **FAILED**
	//
	// 	- **KILLING**
	//
	// 	- **KILLED**
	//
	// 	- **SUCCEEDING**
	//
	// 	- **COMPLETED**
	//
	// 	- **FATAL**
	//
	// 	- **UNKNOWN**
	//
	// example:
	//
	// SUBMITTED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SubmitSparkAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitSparkAppResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *SubmitSparkAppResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *SubmitSparkAppResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SubmitSparkAppResponseBodyData) GetState() *string {
	return s.State
}

func (s *SubmitSparkAppResponseBodyData) SetAppId(v string) *SubmitSparkAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) SetAppName(v string) *SubmitSparkAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) SetMessage(v string) *SubmitSparkAppResponseBodyData {
	s.Message = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) SetState(v string) *SubmitSparkAppResponseBodyData {
	s.State = &v
	return s
}

func (s *SubmitSparkAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
