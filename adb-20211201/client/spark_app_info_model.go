// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkAppInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SparkAppInfo
	GetAppId() *string
	SetAppName(v string) *SparkAppInfo
	GetAppName() *string
	SetDBClusterId(v string) *SparkAppInfo
	GetDBClusterId() *string
	SetDetail(v *Detail) *SparkAppInfo
	GetDetail() *Detail
	SetMessage(v string) *SparkAppInfo
	GetMessage() *string
	SetPriority(v string) *SparkAppInfo
	GetPriority() *string
	SetState(v string) *SparkAppInfo
	GetState() *string
}

type SparkAppInfo struct {
	// example:
	//
	// s202207151211hz0cb4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// Spark Test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// amv-23xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Detail      *Detail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// WARN: Disk is full
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// NORMAL
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkAppInfo) String() string {
	return dara.Prettify(s)
}

func (s SparkAppInfo) GoString() string {
	return s.String()
}

func (s *SparkAppInfo) GetAppId() *string {
	return s.AppId
}

func (s *SparkAppInfo) GetAppName() *string {
	return s.AppName
}

func (s *SparkAppInfo) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SparkAppInfo) GetDetail() *Detail {
	return s.Detail
}

func (s *SparkAppInfo) GetMessage() *string {
	return s.Message
}

func (s *SparkAppInfo) GetPriority() *string {
	return s.Priority
}

func (s *SparkAppInfo) GetState() *string {
	return s.State
}

func (s *SparkAppInfo) SetAppId(v string) *SparkAppInfo {
	s.AppId = &v
	return s
}

func (s *SparkAppInfo) SetAppName(v string) *SparkAppInfo {
	s.AppName = &v
	return s
}

func (s *SparkAppInfo) SetDBClusterId(v string) *SparkAppInfo {
	s.DBClusterId = &v
	return s
}

func (s *SparkAppInfo) SetDetail(v *Detail) *SparkAppInfo {
	s.Detail = v
	return s
}

func (s *SparkAppInfo) SetMessage(v string) *SparkAppInfo {
	s.Message = &v
	return s
}

func (s *SparkAppInfo) SetPriority(v string) *SparkAppInfo {
	s.Priority = &v
	return s
}

func (s *SparkAppInfo) SetState(v string) *SparkAppInfo {
	s.State = &v
	return s
}

func (s *SparkAppInfo) Validate() error {
	return dara.Validate(s)
}
