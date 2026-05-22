// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteCurrentNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNSList(v []*string) *GetSiteCurrentNSResponseBody
	GetNSList() []*string
	SetRequestId(v string) *GetSiteCurrentNSResponseBody
	GetRequestId() *string
}

type GetSiteCurrentNSResponseBody struct {
	NSList    []*string `json:"NSList,omitempty" xml:"NSList,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSiteCurrentNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteCurrentNSResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteCurrentNSResponseBody) GetNSList() []*string {
	return s.NSList
}

func (s *GetSiteCurrentNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteCurrentNSResponseBody) SetNSList(v []*string) *GetSiteCurrentNSResponseBody {
	s.NSList = v
	return s
}

func (s *GetSiteCurrentNSResponseBody) SetRequestId(v string) *GetSiteCurrentNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteCurrentNSResponseBody) Validate() error {
	return dara.Validate(s)
}
