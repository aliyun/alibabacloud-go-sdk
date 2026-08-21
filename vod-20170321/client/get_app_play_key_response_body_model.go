// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPlayKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPlayKey(v *GetAppPlayKeyResponseBodyAppPlayKey) *GetAppPlayKeyResponseBody
	GetAppPlayKey() *GetAppPlayKeyResponseBodyAppPlayKey
	SetRequestId(v string) *GetAppPlayKeyResponseBody
	GetRequestId() *string
}

type GetAppPlayKeyResponseBody struct {
	// The application playback key information.
	AppPlayKey *GetAppPlayKeyResponseBodyAppPlayKey `json:"AppPlayKey,omitempty" xml:"AppPlayKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppPlayKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppPlayKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppPlayKeyResponseBody) GetAppPlayKey() *GetAppPlayKeyResponseBodyAppPlayKey {
	return s.AppPlayKey
}

func (s *GetAppPlayKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppPlayKeyResponseBody) SetAppPlayKey(v *GetAppPlayKeyResponseBodyAppPlayKey) *GetAppPlayKeyResponseBody {
	s.AppPlayKey = v
	return s
}

func (s *GetAppPlayKeyResponseBody) SetRequestId(v string) *GetAppPlayKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppPlayKeyResponseBody) Validate() error {
	if s.AppPlayKey != nil {
		if err := s.AppPlayKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppPlayKeyResponseBodyAppPlayKey struct {
	// The application ID.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the playback key was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2025-03-18T03:59:01Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The time when the playback key was last modified. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2025-03-18T03:59:01Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The playback key.
	//
	// - Only uppercase letters, lowercase letters, and digits are supported. The length must be 8 to 20 characters.
	//
	// - UTF-8 encoding.
	//
	// example:
	//
	// yzNgTUtAl6HAuosIA
	PlayKey *string `json:"PlayKey,omitempty" xml:"PlayKey,omitempty"`
}

func (s GetAppPlayKeyResponseBodyAppPlayKey) String() string {
	return dara.Prettify(s)
}

func (s GetAppPlayKeyResponseBodyAppPlayKey) GoString() string {
	return s.String()
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) GetAppId() *string {
	return s.AppId
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) GetPlayKey() *string {
	return s.PlayKey
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) SetAppId(v string) *GetAppPlayKeyResponseBodyAppPlayKey {
	s.AppId = &v
	return s
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) SetCreationTime(v string) *GetAppPlayKeyResponseBodyAppPlayKey {
	s.CreationTime = &v
	return s
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) SetModificationTime(v string) *GetAppPlayKeyResponseBodyAppPlayKey {
	s.ModificationTime = &v
	return s
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) SetPlayKey(v string) *GetAppPlayKeyResponseBodyAppPlayKey {
	s.PlayKey = &v
	return s
}

func (s *GetAppPlayKeyResponseBodyAppPlayKey) Validate() error {
	return dara.Validate(s)
}
