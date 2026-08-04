// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTrueNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProfileInfo(v *QueryAccountTrueNameResponseBodyProfileInfo) *QueryAccountTrueNameResponseBody
	GetProfileInfo() *QueryAccountTrueNameResponseBodyProfileInfo
	SetRequestId(v string) *QueryAccountTrueNameResponseBody
	GetRequestId() *string
}

type QueryAccountTrueNameResponseBody struct {
	ProfileInfo *QueryAccountTrueNameResponseBodyProfileInfo `json:"ProfileInfo,omitempty" xml:"ProfileInfo,omitempty" type:"Struct"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAccountTrueNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTrueNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountTrueNameResponseBody) GetProfileInfo() *QueryAccountTrueNameResponseBodyProfileInfo {
	return s.ProfileInfo
}

func (s *QueryAccountTrueNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountTrueNameResponseBody) SetProfileInfo(v *QueryAccountTrueNameResponseBodyProfileInfo) *QueryAccountTrueNameResponseBody {
	s.ProfileInfo = v
	return s
}

func (s *QueryAccountTrueNameResponseBody) SetRequestId(v string) *QueryAccountTrueNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountTrueNameResponseBody) Validate() error {
	if s.ProfileInfo != nil {
		if err := s.ProfileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountTrueNameResponseBodyProfileInfo struct {
	TrueName *string `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
}

func (s QueryAccountTrueNameResponseBodyProfileInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTrueNameResponseBodyProfileInfo) GoString() string {
	return s.String()
}

func (s *QueryAccountTrueNameResponseBodyProfileInfo) GetTrueName() *string {
	return s.TrueName
}

func (s *QueryAccountTrueNameResponseBodyProfileInfo) SetTrueName(v string) *QueryAccountTrueNameResponseBodyProfileInfo {
	s.TrueName = &v
	return s
}

func (s *QueryAccountTrueNameResponseBodyProfileInfo) Validate() error {
	return dara.Validate(s)
}
