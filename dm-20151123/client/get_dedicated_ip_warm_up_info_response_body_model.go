// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDedicatedIpWarmUpInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInfo(v []*GetDedicatedIpWarmUpInfoResponseBodyInfo) *GetDedicatedIpWarmUpInfoResponseBody
	GetInfo() []*GetDedicatedIpWarmUpInfoResponseBodyInfo
	SetRequestId(v string) *GetDedicatedIpWarmUpInfoResponseBody
	GetRequestId() *string
}

type GetDedicatedIpWarmUpInfoResponseBody struct {
	Info      []*GetDedicatedIpWarmUpInfoResponseBodyInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDedicatedIpWarmUpInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpInfoResponseBody) GetInfo() []*GetDedicatedIpWarmUpInfoResponseBodyInfo {
	return s.Info
}

func (s *GetDedicatedIpWarmUpInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDedicatedIpWarmUpInfoResponseBody) SetInfo(v []*GetDedicatedIpWarmUpInfoResponseBodyInfo) *GetDedicatedIpWarmUpInfoResponseBody {
	s.Info = v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponseBody) SetRequestId(v string) *GetDedicatedIpWarmUpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDedicatedIpWarmUpInfoResponseBodyInfo struct {
	Esp      *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	Finished *bool   `json:"Finished,omitempty" xml:"Finished,omitempty"`
}

func (s GetDedicatedIpWarmUpInfoResponseBodyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpInfoResponseBodyInfo) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpInfoResponseBodyInfo) GetEsp() *string {
	return s.Esp
}

func (s *GetDedicatedIpWarmUpInfoResponseBodyInfo) GetFinished() *bool {
	return s.Finished
}

func (s *GetDedicatedIpWarmUpInfoResponseBodyInfo) SetEsp(v string) *GetDedicatedIpWarmUpInfoResponseBodyInfo {
	s.Esp = &v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponseBodyInfo) SetFinished(v bool) *GetDedicatedIpWarmUpInfoResponseBodyInfo {
	s.Finished = &v
	return s
}

func (s *GetDedicatedIpWarmUpInfoResponseBodyInfo) Validate() error {
	return dara.Validate(s)
}
