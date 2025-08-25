// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeSuperResolutionImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MakeSuperResolutionImageResponseBodyData) *MakeSuperResolutionImageResponseBody
	GetData() *MakeSuperResolutionImageResponseBodyData
	SetRequestId(v string) *MakeSuperResolutionImageResponseBody
	GetRequestId() *string
}

type MakeSuperResolutionImageResponseBody struct {
	Data *MakeSuperResolutionImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 47DD87F1-D077-499A-8D96-C82F006A6839
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MakeSuperResolutionImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBody) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBody) GetData() *MakeSuperResolutionImageResponseBodyData {
	return s.Data
}

func (s *MakeSuperResolutionImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MakeSuperResolutionImageResponseBody) SetData(v *MakeSuperResolutionImageResponseBodyData) *MakeSuperResolutionImageResponseBody {
	s.Data = v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) SetRequestId(v string) *MakeSuperResolutionImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeSuperResolutionImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type MakeSuperResolutionImageResponseBodyData struct {
	// example:
	//
	// http://ivpd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/upload/ai-gateway_prod/ds%253D20191121/sisrx2_157433961551387538.jpg?Expires=1574598816&OSSAccessKeyId=LTAI4FeJ8qKkYn6SrHhQ****&Signature=8phY6dOz4U889nHfHC1g51nwAi****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MakeSuperResolutionImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *MakeSuperResolutionImageResponseBodyData) SetUrl(v string) *MakeSuperResolutionImageResponseBodyData {
	s.Url = &v
	return s
}

func (s *MakeSuperResolutionImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
