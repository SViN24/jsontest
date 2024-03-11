package main

import "time"

// AUTOGENERATED
type Response struct {
	Success bool `json:"success"`
	Data    []struct {
		Event struct {
			ID                  int         `json:"id"`
			Priority            int         `json:"priority"`
			Status              int         `json:"status"`
			Published           int         `json:"published"`
			ReleaseDateFiles    string      `json:"release_date_files"`
			Expiration          string      `json:"expiration"`
			Title               string      `json:"title"`
			Subtitle            string      `json:"subtitle"`
			Header              interface{} `json:"header"`
			Hours               string      `json:"hours"`
			AbsencesLimit       int         `json:"absences_limit"`
			Enroll              int         `json:"enroll"`
			Index               int         `json:"index"`
			Feed                int         `json:"feed"`
			CertificateTitle    string      `json:"certificate_title"`
			FbGroup             string      `json:"fb_group"`
			EvaluateTopics      interface{} `json:"evaluate_topics"`
			EvaluateInstructors string      `json:"evaluate_instructors"`
			FbTestimonial       string      `json:"fb_testimonial"`
			AuthorID            interface{} `json:"author_id"`
			CreatorID           interface{} `json:"creator_id"`
			ViewTpl             string      `json:"view_tpl"`
			ViewCounter         interface{} `json:"view_counter"`
			PublishedAt         string      `json:"published_at"`
			CreatedAt           time.Time   `json:"created_at"`
			UpdatedAt           time.Time   `json:"updated_at"`
			LaunchDate          string      `json:"launch_date"`
			XMLTitle            string      `json:"xml_title"`
			XMLDescription      string      `json:"xml_description"`
			XMLShortDescription string      `json:"xml_short_description"`
			Pivot               struct {
				UserID        int    `json:"user_id"`
				EventID       int    `json:"event_id"`
				Paid          int    `json:"paid"`
				Expiration    string `json:"expiration"`
				Comment       string `json:"comment"`
				PaymentMethod int    `json:"payment_method"`
			} `json:"pivot"`
			Instructors []struct {
				ID            int         `json:"id"`
				Priority      int         `json:"priority"`
				Status        int         `json:"status"`
				CommentStatus int         `json:"comment_status"`
				Title         string      `json:"title"`
				ShortTitle    string      `json:"short_title"`
				Subtitle      string      `json:"subtitle"`
				Header        string      `json:"header"`
				Company       string      `json:"company"`
				Summary       string      `json:"summary"`
				Body          string      `json:"body"`
				ExtURL        string      `json:"ext_url"`
				SocialMedia   string      `json:"social_media"`
				AuthorID      interface{} `json:"author_id"`
				CreatorID     interface{} `json:"creator_id"`
				CreatedAt     time.Time   `json:"created_at"`
				UpdatedAt     time.Time   `json:"updated_at"`
				LessonID      int         `json:"lesson_id"`
				InstructorID  int         `json:"instructor_id"`
				EventID       int         `json:"event_id"`
				Pivot         struct {
					EventID      int `json:"event_id"`
					InstructorID int `json:"instructor_id"`
					LessonID     int `json:"lesson_id"`
				} `json:"pivot"`
			} `json:"instructors"`
			EventInfo1 struct {
				ID                             int         `json:"id"`
				EventID                        int         `json:"event_id"`
				CourseStatus                   string      `json:"course_status"`
				CourseHours                    string      `json:"course_hours"`
				CourseHoursText                string      `json:"course_hours_text"`
				CourseHoursVisible             string      `json:"course_hours_visible"`
				CourseHoursIcon                string      `json:"course_hours_icon"`
				CourseLanguage                 string      `json:"course_language"`
				CourseLanguageVisible          string      `json:"course_language_visible"`
				CourseLanguageIcon             string      `json:"course_language_icon"`
				CourseDelivery                 string      `json:"course_delivery"`
				CourseElearningIcon            string      `json:"course_elearning_icon"`
				CourseElearningVisible         string      `json:"course_elearning_visible"`
				CourseElearningExpiration      int         `json:"course_elearning_expiration"`
				CourseElearningText            string      `json:"course_elearning_text"`
				CourseElearningExamIcon        string      `json:"course_elearning_exam_icon"`
				CourseElearningExamVisible     string      `json:"course_elearning_exam_visible"`
				CourseElearningExamText        interface{} `json:"course_elearning_exam_text"`
				CourseInclassCity              interface{} `json:"course_inclass_city"`
				CourseInclassCityIcon          string      `json:"course_inclass_city_icon"`
				CourseInclassDates             string      `json:"course_inclass_dates"`
				CourseInclassDays              string      `json:"course_inclass_days"`
				CourseInclassTimes             string      `json:"course_inclass_times"`
				CourseInclassAbsences          interface{} `json:"course_inclass_absences"`
				CourseElearningAccess          interface{} `json:"course_elearning_access"`
				CourseElearningAccessIcon      string      `json:"course_elearning_access_icon"`
				CoursePaymentMethod            string      `json:"course_payment_method"`
				CoursePaymentIcon              string      `json:"course_payment_icon"`
				CourseFilesIcon                string      `json:"course_files_icon"`
				CoursePartner                  int         `json:"course_partner"`
				CoursePartnerIcon              string      `json:"course_partner_icon"`
				CourseManager                  int         `json:"course_manager"`
				CourseManagerIcon              string      `json:"course_manager_icon"`
				CourseAwards                   int         `json:"course_awards"`
				CourseAwardsText               interface{} `json:"course_awards_text"`
				CourseAwardsIcon               string      `json:"course_awards_icon"`
				HasCertificate                 int         `json:"has_certificate"`
				CourseCertificationNameSuccess string      `json:"course_certification_name_success"`
				CourseCertificationNameFailure string      `json:"course_certification_name_failure"`
				CourseCertificationEventTitle  string      `json:"course_certification_event_title"`
				CourseCertificationType        string      `json:"course_certification_type"`
				CourseCertificationVisible     string      `json:"course_certification_visible"`
				CourseCertificationIcon        string      `json:"course_certification_icon"`
				CourseStudentsNumber           string      `json:"course_students_number"`
				CourseStudentsText             string      `json:"course_students_text"`
				CourseStudentsVisible          string      `json:"course_students_visible"`
				CourseStudentsIcon             string      `json:"course_students_icon"`
				CreatedAt                      time.Time   `json:"created_at"`
				UpdatedAt                      time.Time   `json:"updated_at"`
			} `json:"event_info1"`
		} `json:"event"`
		Summary []struct {
			Title       interface{} `json:"title"`
			Description string      `json:"description"`
			Icon        interface{} `json:"icon"`
			Section     string      `json:"section"`
		} `json:"summary"`
		IsElearning   bool    `json:"is_elearning,omitempty"`
		Progress      string  `json:"progress,omitempty"`
		VideosSeen    string  `json:"videos_seen,omitempty"`
		LastVideoSeen int     `json:"lastVideoSeen,omitempty"`
		Topics        []Topic `json:"topics"`
		IsInclass     bool    `json:"is_inclass,omitempty"`
		Date          string  `json:"date,omitempty"`
		Files         struct {
			Folders []interface{} `json:"folders"` //empty
		} `json:"files,omitempty"`
		Calendar []struct {
			Time            string `json:"time"`
			DateTime        string `json:"date_time"`
			Title           string `json:"title"`
			Room            string `json:"room"`
			InstructorImage string `json:"instructor_image"`
			InstructorName  string `json:"instructor_name"`
		} `json:"calendar,omitempty"`
	} `json:"data"`
}

// autogenerated 2
type Topic []struct {
	TopicContent struct {
		Lessons []struct {
			Date       string `json:"date"`
			Title      string `json:"title"`
			TimeStarts string `json:"time_starts"`
			TimeEnds   string `json:"time_ends"`
			Duration   string `json:"duration"`
			Room       string `json:"room"`
			Instructor struct {
				Name  string `json:"name"`
				Media string `json:"media"`
			} `json:"instructor"`
		} `json:"lessons"`
		TotalDuration string `json:"total_duration"`
		TopicID       int    `json:"topic_id"`
	} `json:"topic_content"`
	TopicName string `json:"topic_name"`
}
