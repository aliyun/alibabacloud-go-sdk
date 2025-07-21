// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOssUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectNameList(v []*string) *GenerateOssUrlRequest
	GetObjectNameList() []*string
	SetSessionId(v string) *GenerateOssUrlRequest
	GetSessionId() *string
}

type GenerateOssUrlRequest struct {
	ObjectNameList []*string `json:"ObjectNameList,omitempty" xml:"ObjectNameList,omitempty" type:"Repeated"`
	SessionId      *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GenerateOssUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateOssUrlRequest) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlRequest) GetObjectNameList() []*string {
	return s.ObjectNameList
}

func (s *GenerateOssUrlRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenerateOssUrlRequest) SetObjectNameList(v []*string) *GenerateOssUrlRequest {
	s.ObjectNameList = v
	return s
}

func (s *GenerateOssUrlRequest) SetSessionId(v string) *GenerateOssUrlRequest {
	s.SessionId = &v
	return s
}

func (s *GenerateOssUrlRequest) Validate() error {
	return dara.Validate(s)
}
