// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSFlowStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOSSFlowStatisList(v []*GetOSSFlowStatisResponseBodyOSSFlowStatisList) *GetOSSFlowStatisResponseBody
	GetOSSFlowStatisList() []*GetOSSFlowStatisResponseBodyOSSFlowStatisList
	SetRequestId(v string) *GetOSSFlowStatisResponseBody
	GetRequestId() *string
}

type GetOSSFlowStatisResponseBody struct {
	OSSFlowStatisList []*GetOSSFlowStatisResponseBodyOSSFlowStatisList `json:"OSSFlowStatisList,omitempty" xml:"OSSFlowStatisList,omitempty" type:"Repeated"`
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOSSFlowStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOSSFlowStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSFlowStatisResponseBody) GetOSSFlowStatisList() []*GetOSSFlowStatisResponseBodyOSSFlowStatisList {
	return s.OSSFlowStatisList
}

func (s *GetOSSFlowStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOSSFlowStatisResponseBody) SetOSSFlowStatisList(v []*GetOSSFlowStatisResponseBodyOSSFlowStatisList) *GetOSSFlowStatisResponseBody {
	s.OSSFlowStatisList = v
	return s
}

func (s *GetOSSFlowStatisResponseBody) SetRequestId(v string) *GetOSSFlowStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOSSFlowStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOSSFlowStatisResponseBodyOSSFlowStatisList struct {
	NetworkOut  *int64  `json:"NetworkOut,omitempty" xml:"NetworkOut,omitempty"`
	StatTime    *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
	StatTimeUTC *string `json:"StatTimeUTC,omitempty" xml:"StatTimeUTC,omitempty"`
}

func (s GetOSSFlowStatisResponseBodyOSSFlowStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetOSSFlowStatisResponseBodyOSSFlowStatisList) GoString() string {
	return s.String()
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) GetNetworkOut() *int64 {
	return s.NetworkOut
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) GetStatTime() *string {
	return s.StatTime
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) GetStatTimeUTC() *string {
	return s.StatTimeUTC
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) SetNetworkOut(v int64) *GetOSSFlowStatisResponseBodyOSSFlowStatisList {
	s.NetworkOut = &v
	return s
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) SetStatTime(v string) *GetOSSFlowStatisResponseBodyOSSFlowStatisList {
	s.StatTime = &v
	return s
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) SetStatTimeUTC(v string) *GetOSSFlowStatisResponseBodyOSSFlowStatisList {
	s.StatTimeUTC = &v
	return s
}

func (s *GetOSSFlowStatisResponseBodyOSSFlowStatisList) Validate() error {
	return dara.Validate(s)
}
