// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrustedProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *GetTrustedProjectsResponseBody
	GetData() []*string
	SetRequestId(v string) *GetTrustedProjectsResponseBody
	GetRequestId() *string
}

type GetTrustedProjectsResponseBody struct {
	// The returned data.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc0590416675329272834336e4387
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTrustedProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrustedProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustedProjectsResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetTrustedProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrustedProjectsResponseBody) SetData(v []*string) *GetTrustedProjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetTrustedProjectsResponseBody) SetRequestId(v string) *GetTrustedProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrustedProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
