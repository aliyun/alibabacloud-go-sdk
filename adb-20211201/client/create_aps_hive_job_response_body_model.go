// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsHiveJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApsHiveJobResponseBody
	GetCode() *string
	SetData(v string) *CreateApsHiveJobResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *CreateApsHiveJobResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateApsHiveJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApsHiveJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApsHiveJobResponseBody
	GetSuccess() *bool
}

type CreateApsHiveJobResponseBody struct {
	// The response code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2895BB82-B2C1-408E-AA73-DB8D59******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApsHiveJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApsHiveJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApsHiveJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApsHiveJobResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApsHiveJobResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateApsHiveJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApsHiveJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApsHiveJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApsHiveJobResponseBody) SetCode(v string) *CreateApsHiveJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApsHiveJobResponseBody) SetData(v string) *CreateApsHiveJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApsHiveJobResponseBody) SetHttpStatusCode(v string) *CreateApsHiveJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApsHiveJobResponseBody) SetMessage(v string) *CreateApsHiveJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApsHiveJobResponseBody) SetRequestId(v string) *CreateApsHiveJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApsHiveJobResponseBody) SetSuccess(v bool) *CreateApsHiveJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApsHiveJobResponseBody) Validate() error {
	return dara.Validate(s)
}
