// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMusicTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryMusicTypeResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryMusicTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMusicTypeResponseBody
	GetRequestId() *string
	SetResult(v []*QueryMusicTypeResponseBodyResult) *QueryMusicTypeResponseBody
	GetResult() []*QueryMusicTypeResponseBodyResult
}

type QueryMusicTypeResponseBody struct {
	// Status code returned by the alarm service
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// 设备账号未关联
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// request ID
	//
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of ringtone types
	Result []*QueryMusicTypeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryMusicTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryMusicTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMusicTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMusicTypeResponseBody) GetResult() []*QueryMusicTypeResponseBodyResult {
	return s.Result
}

func (s *QueryMusicTypeResponseBody) SetCode(v int32) *QueryMusicTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMusicTypeResponseBody) SetMessage(v string) *QueryMusicTypeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMusicTypeResponseBody) SetRequestId(v string) *QueryMusicTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMusicTypeResponseBody) SetResult(v []*QueryMusicTypeResponseBodyResult) *QueryMusicTypeResponseBody {
	s.Result = v
	return s
}

func (s *QueryMusicTypeResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMusicTypeResponseBodyResult struct {
	// Ringtone type ID
	//
	// example:
	//
	// 1
	MusicType *int64 `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// Name of the ringtone category
	//
	// example:
	//
	// xx
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
}

func (s QueryMusicTypeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeResponseBodyResult) GetMusicType() *int64 {
	return s.MusicType
}

func (s *QueryMusicTypeResponseBodyResult) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *QueryMusicTypeResponseBodyResult) SetMusicType(v int64) *QueryMusicTypeResponseBodyResult {
	s.MusicType = &v
	return s
}

func (s *QueryMusicTypeResponseBodyResult) SetMusicTypeName(v string) *QueryMusicTypeResponseBodyResult {
	s.MusicTypeName = &v
	return s
}

func (s *QueryMusicTypeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
