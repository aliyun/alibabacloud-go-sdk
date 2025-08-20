// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsDatasoureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApsDatasoureResponseBody
	GetCode() *string
	SetDBClusterId(v string) *CreateApsDatasoureResponseBody
	GetDBClusterId() *string
	SetData(v string) *CreateApsDatasoureResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApsDatasoureResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApsDatasoureResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApsDatasoureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApsDatasoureResponseBody
	GetSuccess() *bool
}

type CreateApsDatasoureResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-7xxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the cluster resource usage.
	//
	// example:
	//
	// 69
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65D2***-45C1-5C18-**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the dry run succeeds. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApsDatasoureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApsDatasoureResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsDatasoureResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApsDatasoureResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApsDatasoureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApsDatasoureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApsDatasoureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApsDatasoureResponseBody) SetCode(v string) *CreateApsDatasoureResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) SetDBClusterId(v string) *CreateApsDatasoureResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) SetData(v string) *CreateApsDatasoureResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) SetHttpStatusCode(v int32) *CreateApsDatasoureResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) SetMessage(v string) *CreateApsDatasoureResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) SetRequestId(v string) *CreateApsDatasoureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) SetSuccess(v bool) *CreateApsDatasoureResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApsDatasoureResponseBody) Validate() error {
	return dara.Validate(s)
}
