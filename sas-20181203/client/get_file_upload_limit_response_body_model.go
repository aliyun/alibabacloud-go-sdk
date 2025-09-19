// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileUploadLimitResponseBodyData) *GetFileUploadLimitResponseBody
	GetData() *GetFileUploadLimitResponseBodyData
	SetRequestId(v string) *GetFileUploadLimitResponseBody
	GetRequestId() *string
}

type GetFileUploadLimitResponseBody struct {
	// The data returned if the request is successful.
	Data *GetFileUploadLimitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1A975D03-5F49-5354-B2CB-3918D5DA431A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileUploadLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadLimitResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileUploadLimitResponseBody) GetData() *GetFileUploadLimitResponseBodyData {
	return s.Data
}

func (s *GetFileUploadLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileUploadLimitResponseBody) SetData(v *GetFileUploadLimitResponseBodyData) *GetFileUploadLimitResponseBody {
	s.Data = v
	return s
}

func (s *GetFileUploadLimitResponseBody) SetRequestId(v string) *GetFileUploadLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileUploadLimitResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileUploadLimitResponseBodyData struct {
	// The QPS limit on the files uploaded from the client.
	//
	// example:
	//
	// 100
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s GetFileUploadLimitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadLimitResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileUploadLimitResponseBodyData) GetLimit() *string {
	return s.Limit
}

func (s *GetFileUploadLimitResponseBodyData) SetLimit(v string) *GetFileUploadLimitResponseBodyData {
	s.Limit = &v
	return s
}

func (s *GetFileUploadLimitResponseBodyData) Validate() error {
	return dara.Validate(s)
}
