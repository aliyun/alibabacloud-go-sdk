// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogMeta(v *GetLogMetaResponseBodyLogMeta) *GetLogMetaResponseBody
	GetLogMeta() *GetLogMetaResponseBodyLogMeta
	SetRequestId(v string) *GetLogMetaResponseBody
	GetRequestId() *string
}

type GetLogMetaResponseBody struct {
	// The data of a data shipping task.
	LogMeta *GetLogMetaResponseBodyLogMeta `json:"LogMeta,omitempty" xml:"LogMeta,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLogMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogMetaResponseBody) GetLogMeta() *GetLogMetaResponseBodyLogMeta {
	return s.LogMeta
}

func (s *GetLogMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogMetaResponseBody) SetLogMeta(v *GetLogMetaResponseBodyLogMeta) *GetLogMetaResponseBody {
	s.LogMeta = v
	return s
}

func (s *GetLogMetaResponseBody) SetRequestId(v string) *GetLogMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogMetaResponseBody) Validate() error {
	if s.LogMeta != nil {
		if err := s.LogMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLogMetaResponseBodyLogMeta struct {
	// The name of the dedicated Logstore in which logs are stored.
	//
	// example:
	//
	// aegis-log-login
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// sas-log
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The status of a data shipping task of a log. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLogMetaResponseBodyLogMeta) String() string {
	return dara.Prettify(s)
}

func (s GetLogMetaResponseBodyLogMeta) GoString() string {
	return s.String()
}

func (s *GetLogMetaResponseBodyLogMeta) GetLogStore() *string {
	return s.LogStore
}

func (s *GetLogMetaResponseBodyLogMeta) GetProject() *string {
	return s.Project
}

func (s *GetLogMetaResponseBodyLogMeta) GetStatus() *string {
	return s.Status
}

func (s *GetLogMetaResponseBodyLogMeta) SetLogStore(v string) *GetLogMetaResponseBodyLogMeta {
	s.LogStore = &v
	return s
}

func (s *GetLogMetaResponseBodyLogMeta) SetProject(v string) *GetLogMetaResponseBodyLogMeta {
	s.Project = &v
	return s
}

func (s *GetLogMetaResponseBodyLogMeta) SetStatus(v string) *GetLogMetaResponseBodyLogMeta {
	s.Status = &v
	return s
}

func (s *GetLogMetaResponseBodyLogMeta) Validate() error {
	return dara.Validate(s)
}
