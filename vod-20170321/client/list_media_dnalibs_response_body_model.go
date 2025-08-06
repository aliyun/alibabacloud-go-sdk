// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaDNALibsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIDNALibInfoList(v *ListMediaDNALibsResponseBodyAIDNALibInfoList) *ListMediaDNALibsResponseBody
	GetAIDNALibInfoList() *ListMediaDNALibsResponseBodyAIDNALibInfoList
	SetRequestId(v string) *ListMediaDNALibsResponseBody
	GetRequestId() *string
}

type ListMediaDNALibsResponseBody struct {
	AIDNALibInfoList *ListMediaDNALibsResponseBodyAIDNALibInfoList `json:"AIDNALibInfoList,omitempty" xml:"AIDNALibInfoList,omitempty" type:"Struct"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaDNALibsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNALibsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaDNALibsResponseBody) GetAIDNALibInfoList() *ListMediaDNALibsResponseBodyAIDNALibInfoList {
	return s.AIDNALibInfoList
}

func (s *ListMediaDNALibsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMediaDNALibsResponseBody) SetAIDNALibInfoList(v *ListMediaDNALibsResponseBodyAIDNALibInfoList) *ListMediaDNALibsResponseBody {
	s.AIDNALibInfoList = v
	return s
}

func (s *ListMediaDNALibsResponseBody) SetRequestId(v string) *ListMediaDNALibsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaDNALibsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMediaDNALibsResponseBodyAIDNALibInfoList struct {
	AIDNALibInfo []*ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo `json:"AIDNALibInfo,omitempty" xml:"AIDNALibInfo,omitempty" type:"Repeated"`
}

func (s ListMediaDNALibsResponseBodyAIDNALibInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNALibsResponseBodyAIDNALibInfoList) GoString() string {
	return s.String()
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoList) GetAIDNALibInfo() []*ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo {
	return s.AIDNALibInfo
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoList) SetAIDNALibInfo(v []*ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) *ListMediaDNALibsResponseBodyAIDNALibInfoList {
	s.AIDNALibInfo = v
	return s
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoList) Validate() error {
	return dara.Validate(s)
}

type ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo struct {
	FpDBId    *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) GoString() string {
	return s.String()
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) GetFpDBId() *string {
	return s.FpDBId
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) GetModelType() *string {
	return s.ModelType
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) GetState() *string {
	return s.State
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) SetFpDBId(v string) *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo {
	s.FpDBId = &v
	return s
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) SetModelType(v string) *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo {
	s.ModelType = &v
	return s
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) SetState(v string) *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo {
	s.State = &v
	return s
}

func (s *ListMediaDNALibsResponseBodyAIDNALibInfoListAIDNALibInfo) Validate() error {
	return dara.Validate(s)
}
