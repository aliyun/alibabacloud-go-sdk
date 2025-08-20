// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsDatasoureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApsDatasoureResponseBody
	GetCode() *string
	SetData(v string) *DeleteApsDatasoureResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *DeleteApsDatasoureResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DeleteApsDatasoureResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApsDatasoureResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteApsDatasoureResponseBody
	GetSuccess() *string
}

type DeleteApsDatasoureResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 7
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// FAE98A4F-****-****-BF6D-67EEAC9C39DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApsDatasoureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsDatasoureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApsDatasoureResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApsDatasoureResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteApsDatasoureResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteApsDatasoureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApsDatasoureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApsDatasoureResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteApsDatasoureResponseBody) SetCode(v string) *DeleteApsDatasoureResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApsDatasoureResponseBody) SetData(v string) *DeleteApsDatasoureResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteApsDatasoureResponseBody) SetHttpStatusCode(v string) *DeleteApsDatasoureResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteApsDatasoureResponseBody) SetMessage(v string) *DeleteApsDatasoureResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApsDatasoureResponseBody) SetRequestId(v string) *DeleteApsDatasoureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApsDatasoureResponseBody) SetSuccess(v string) *DeleteApsDatasoureResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApsDatasoureResponseBody) Validate() error {
	return dara.Validate(s)
}
