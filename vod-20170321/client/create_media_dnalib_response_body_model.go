// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaDNALibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIDNALibInfo(v *CreateMediaDNALibResponseBodyAIDNALibInfo) *CreateMediaDNALibResponseBody
	GetAIDNALibInfo() *CreateMediaDNALibResponseBodyAIDNALibInfo
	SetRequestId(v string) *CreateMediaDNALibResponseBody
	GetRequestId() *string
}

type CreateMediaDNALibResponseBody struct {
	AIDNALibInfo *CreateMediaDNALibResponseBodyAIDNALibInfo `json:"AIDNALibInfo,omitempty" xml:"AIDNALibInfo,omitempty" type:"Struct"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMediaDNALibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaDNALibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaDNALibResponseBody) GetAIDNALibInfo() *CreateMediaDNALibResponseBodyAIDNALibInfo {
	return s.AIDNALibInfo
}

func (s *CreateMediaDNALibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaDNALibResponseBody) SetAIDNALibInfo(v *CreateMediaDNALibResponseBodyAIDNALibInfo) *CreateMediaDNALibResponseBody {
	s.AIDNALibInfo = v
	return s
}

func (s *CreateMediaDNALibResponseBody) SetRequestId(v string) *CreateMediaDNALibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaDNALibResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMediaDNALibResponseBodyAIDNALibInfo struct {
	FpDBId    *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateMediaDNALibResponseBodyAIDNALibInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaDNALibResponseBodyAIDNALibInfo) GoString() string {
	return s.String()
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) GetFpDBId() *string {
	return s.FpDBId
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) GetModelType() *string {
	return s.ModelType
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) GetState() *string {
	return s.State
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) SetFpDBId(v string) *CreateMediaDNALibResponseBodyAIDNALibInfo {
	s.FpDBId = &v
	return s
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) SetModelType(v string) *CreateMediaDNALibResponseBodyAIDNALibInfo {
	s.ModelType = &v
	return s
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) SetState(v string) *CreateMediaDNALibResponseBodyAIDNALibInfo {
	s.State = &v
	return s
}

func (s *CreateMediaDNALibResponseBodyAIDNALibInfo) Validate() error {
	return dara.Validate(s)
}
