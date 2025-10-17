// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppAttemptLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkAppAttemptLogResponseBodyData) *GetSparkAppAttemptLogResponseBody
	GetData() *GetSparkAppAttemptLogResponseBodyData
	SetRequestId(v string) *GetSparkAppAttemptLogResponseBody
	GetRequestId() *string
}

type GetSparkAppAttemptLogResponseBody struct {
	// The queried log.
	Data *GetSparkAppAttemptLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C3A9594F-1D40-4472-A96C-8FB8AA20D38C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppAttemptLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppAttemptLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogResponseBody) GetData() *GetSparkAppAttemptLogResponseBodyData {
	return s.Data
}

func (s *GetSparkAppAttemptLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkAppAttemptLogResponseBody) SetData(v *GetSparkAppAttemptLogResponseBodyData) *GetSparkAppAttemptLogResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppAttemptLogResponseBody) SetRequestId(v string) *GetSparkAppAttemptLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSparkAppAttemptLogResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// s202204132018hzprec1ac61a000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-clusterxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The content of the log.
	//
	// example:
	//
	// 22/04/22 15:30:49 INFO Utils: Start the dump task because s202207151211hz****-0001 app end, the interval is 238141ms;22/04/22 15:30:49 INFO AbstractConnector: Stopped Spark@5e774d9d{HTTP/1.1, (http/1.1)}{0.0.0.0:4040}
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The number of log entries. A value of 0 indicates that no valid logs are returned.
	//
	// example:
	//
	// 775946240
	LogSize *int32 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The alert message returned for the request, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	//
	// example:
	//
	// WARNING: log file maybe deleted, please check oss path: oss://TestBucketName/applog/
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetSparkAppAttemptLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppAttemptLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppAttemptLogResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppAttemptLogResponseBodyData) GetLogContent() *string {
	return s.LogContent
}

func (s *GetSparkAppAttemptLogResponseBodyData) GetLogSize() *int32 {
	return s.LogSize
}

func (s *GetSparkAppAttemptLogResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetAppId(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetDBClusterId(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetLogContent(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.LogContent = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetLogSize(v int32) *GetSparkAppAttemptLogResponseBodyData {
	s.LogSize = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) SetMessage(v string) *GetSparkAppAttemptLogResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetSparkAppAttemptLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
