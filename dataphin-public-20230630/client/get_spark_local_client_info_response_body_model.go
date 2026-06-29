// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkLocalClientInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSparkLocalClientInfoResponseBody
	GetCode() *string
	SetData(v []*GetSparkLocalClientInfoResponseBodyData) *GetSparkLocalClientInfoResponseBody
	GetData() []*GetSparkLocalClientInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetSparkLocalClientInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSparkLocalClientInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSparkLocalClientInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSparkLocalClientInfoResponseBody
	GetSuccess() *bool
}

type GetSparkLocalClientInfoResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*GetSparkLocalClientInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSparkLocalClientInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLocalClientInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkLocalClientInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSparkLocalClientInfoResponseBody) GetData() []*GetSparkLocalClientInfoResponseBodyData {
	return s.Data
}

func (s *GetSparkLocalClientInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSparkLocalClientInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSparkLocalClientInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkLocalClientInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSparkLocalClientInfoResponseBody) SetCode(v string) *GetSparkLocalClientInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBody) SetData(v []*GetSparkLocalClientInfoResponseBodyData) *GetSparkLocalClientInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkLocalClientInfoResponseBody) SetHttpStatusCode(v int32) *GetSparkLocalClientInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBody) SetMessage(v string) *GetSparkLocalClientInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBody) SetRequestId(v string) *GetSparkLocalClientInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBody) SetSuccess(v bool) *GetSparkLocalClientInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSparkLocalClientInfoResponseBodyData struct {
	// The client file ID.
	//
	// example:
	//
	// XXXX-XXXX-XXXX-XXXX
	ClientFileResourceId *string `json:"ClientFileResourceId,omitempty" xml:"ClientFileResourceId,omitempty"`
	// The client file name.
	//
	// example:
	//
	// spark-cleint.zip
	ClientFileResourceName *string `json:"ClientFileResourceName,omitempty" xml:"ClientFileResourceName,omitempty"`
	// The client name.
	//
	// example:
	//
	// spark-clinet
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// Indicates whether the client information is editable.
	//
	// example:
	//
	// true/false
	Editable *bool `json:"Editable,omitempty" xml:"Editable,omitempty"`
}

func (s GetSparkLocalClientInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLocalClientInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkLocalClientInfoResponseBodyData) GetClientFileResourceId() *string {
	return s.ClientFileResourceId
}

func (s *GetSparkLocalClientInfoResponseBodyData) GetClientFileResourceName() *string {
	return s.ClientFileResourceName
}

func (s *GetSparkLocalClientInfoResponseBodyData) GetClientName() *string {
	return s.ClientName
}

func (s *GetSparkLocalClientInfoResponseBodyData) GetEditable() *bool {
	return s.Editable
}

func (s *GetSparkLocalClientInfoResponseBodyData) SetClientFileResourceId(v string) *GetSparkLocalClientInfoResponseBodyData {
	s.ClientFileResourceId = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBodyData) SetClientFileResourceName(v string) *GetSparkLocalClientInfoResponseBodyData {
	s.ClientFileResourceName = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBodyData) SetClientName(v string) *GetSparkLocalClientInfoResponseBodyData {
	s.ClientName = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBodyData) SetEditable(v bool) *GetSparkLocalClientInfoResponseBodyData {
	s.Editable = &v
	return s
}

func (s *GetSparkLocalClientInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
