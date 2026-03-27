// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudResourceDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v [][]*string) *GetCloudResourceDataResponseBody
	GetData() [][]*string
	SetHeader(v []*string) *GetCloudResourceDataResponseBody
	GetHeader() []*string
	SetRequestId(v string) *GetCloudResourceDataResponseBody
	GetRequestId() *string
}

type GetCloudResourceDataResponseBody struct {
	Data   [][]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Header []*string   `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCloudResourceDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudResourceDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudResourceDataResponseBody) GetData() [][]*string {
	return s.Data
}

func (s *GetCloudResourceDataResponseBody) GetHeader() []*string {
	return s.Header
}

func (s *GetCloudResourceDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudResourceDataResponseBody) SetData(v [][]*string) *GetCloudResourceDataResponseBody {
	s.Data = v
	return s
}

func (s *GetCloudResourceDataResponseBody) SetHeader(v []*string) *GetCloudResourceDataResponseBody {
	s.Header = v
	return s
}

func (s *GetCloudResourceDataResponseBody) SetRequestId(v string) *GetCloudResourceDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudResourceDataResponseBody) Validate() error {
	return dara.Validate(s)
}
