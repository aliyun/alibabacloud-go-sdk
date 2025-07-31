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
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetSparkLocalClientInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	return dara.Validate(s)
}

type GetSparkLocalClientInfoResponseBodyData struct {
	// example:
	//
	// XXXX-XXXX-XXXX-XXXX
	ClientFileResourceId *string `json:"ClientFileResourceId,omitempty" xml:"ClientFileResourceId,omitempty"`
	// example:
	//
	// spark-cleint.zip
	ClientFileResourceName *string `json:"ClientFileResourceName,omitempty" xml:"ClientFileResourceName,omitempty"`
	// example:
	//
	// spark-clinet
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
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
