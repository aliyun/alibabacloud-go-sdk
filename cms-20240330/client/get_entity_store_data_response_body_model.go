// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v [][]*string) *GetEntityStoreDataResponseBody
	GetData() [][]*string
	SetHeader(v []*string) *GetEntityStoreDataResponseBody
	GetHeader() []*string
	SetRequestId(v string) *GetEntityStoreDataResponseBody
	GetRequestId() *string
}

type GetEntityStoreDataResponseBody struct {
	Data   [][]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Header []*string   `json:"header,omitempty" xml:"header,omitempty" type:"Repeated"`
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEntityStoreDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponseBody) GetData() [][]*string {
	return s.Data
}

func (s *GetEntityStoreDataResponseBody) GetHeader() []*string {
	return s.Header
}

func (s *GetEntityStoreDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEntityStoreDataResponseBody) SetData(v [][]*string) *GetEntityStoreDataResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetHeader(v []*string) *GetEntityStoreDataResponseBody {
	s.Header = v
	return s
}

func (s *GetEntityStoreDataResponseBody) SetRequestId(v string) *GetEntityStoreDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityStoreDataResponseBody) Validate() error {
	return dara.Validate(s)
}
