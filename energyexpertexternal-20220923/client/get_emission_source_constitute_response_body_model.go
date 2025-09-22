// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmissionSourceConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ConstituteItem) *GetEmissionSourceConstituteResponseBody
	GetData() []*ConstituteItem
	SetRequestId(v string) *GetEmissionSourceConstituteResponseBody
	GetRequestId() *string
}

type GetEmissionSourceConstituteResponseBody struct {
	// Response parameters
	Data []*ConstituteItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEmissionSourceConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSourceConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmissionSourceConstituteResponseBody) GetData() []*ConstituteItem {
	return s.Data
}

func (s *GetEmissionSourceConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmissionSourceConstituteResponseBody) SetData(v []*ConstituteItem) *GetEmissionSourceConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetEmissionSourceConstituteResponseBody) SetRequestId(v string) *GetEmissionSourceConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmissionSourceConstituteResponseBody) Validate() error {
	return dara.Validate(s)
}
